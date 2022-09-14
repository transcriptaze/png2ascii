#/bin/bash

# apt -yq install ffmpeg
# apt -yq install imagemagick
# 
# mkdir -p /opt/mp42asc
# mkdir -p /opt/mp42asc/frames
# mkdir -p /opt/mp42asc/out
# 
# export PATH=$PATH:/opt/mp42asc
# 
# NOTE: edit memory/size policy in /etc/ImageMagick-6/policy.xml:
# 
# <policy domain="resource" name="memory" value="4GB"/>
#  <policy domain="resource" name="map" value="512MiB"/>
#  <policy domain="resource" name="width" value="32KP"/>
#  <policy domain="resource" name="height" value="32KP"/>
#  <policy domain="resource" name="area" value="1GP"/>
#  <policy domain="resource" name="disk" value="1GiB"/>

FILEID="$1"

convert () {
    filename=$(basename $1)

    echo "${filename}"
    mp42asc -debug -format png -out "./out/${filename}" $1
    mogrify -monitor -resize 1920x1080! "./out/${filename}"
    rm $1
}

export -f convert

cd /opt/mp42asc

COOKIES=$(wget --quiet --save-cookies /tmp/cookies.txt --keep-session-cookies --no-check-certificate "https://docs.google.com/uc?export=download&id=${FILEID}" -O- | sed -rn 's/.*confirm=([0-9A-Za-z_]+).*/\1\n/p')
wget --load-cookies /tmp/cookies.txt "https://docs.google.com/uc?export=download&confirm=${COOKIES}&id=${FILEID}" -O movie.mp4 && rm -rf /tmp/cookies.txt

ffmpeg -i movie.mp4 out.wav
ffmpeg -i movie.mp4 -f image2 -c:v png frames/frame_%05d.png

mogrify -monitor -rotate 90 frames/*.png

find frames -name "*.png" -exec bash -c 'convert "$0"' {} \;

ffmpeg -r 29.92 -f image2 -i out/frame_%05d.png -pix_fmt yuv420p out.mp4
ffmpeg -i out.mp4 -i out.wav -c:v copy -c:a aac mp42asc.mp4

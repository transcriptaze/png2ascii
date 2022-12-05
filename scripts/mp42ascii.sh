#/bin/bash

# apt -yq install ffmpeg
# apt -yq install imagemagick
# apt -yq install golang
# 
# mkdir -p /opt
# cd /opt
# git clone https://github.com/transcriptaze/png2ascii
# cd png2ascii
# make build
# export PATH=$PATH:/opt/png2ascii/bin
#
# mkdir -p /opt/mp42asc
# mkdir -p /opt/mp42asc/frames
# mkdir -p /opt/mp42asc/out
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
MP4="movie.mov"

convert () {
    filename=$(basename $1)

    echo "${filename}"
    png2ascii --debug --profile profile.json --format png --out "./out/${filename}" $1
    mogrify   -monitor -resize 1920x1080^ -crop 1920x1080+0+0 "./out/${filename}"
    rm $1
}

export -f convert

cd /opt/mp42asc

COOKIES=$(wget --quiet --save-cookies /tmp/cookies.txt --keep-session-cookies --no-check-certificate "https://docs.google.com/uc?export=download&id=${FILEID}" -O- | sed -rn 's/.*confirm=([0-9A-Za-z_]+).*/\1\n/p')
wget --load-cookies /tmp/cookies.txt "https://docs.google.com/uc?export=download&confirm=${COOKIES}&id=${FILEID}" -O "${MP4}" && rm -rf /tmp/cookies.txt

ffmpeg -i "${MP4}" out.wav
ffmpeg -i "${MP4}" -f image2 -c:v png frames/frame_%05d.png

# ROTATE
# mogrify -monitor -rotate 90 frames/*.png

find frames -name "*.png" -exec bash -c 'convert "$0"' {} \;

ffmpeg -r 29.92 -f image2 -i out/frame_%05d.png -pix_fmt yuv420p out.mp4
ffmpeg -i out.mp4 -i out.wav -c:v copy -c:a aac mp42asc.mp4

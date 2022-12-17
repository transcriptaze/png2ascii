#/bin/bash

apt -yq install ffmpeg
apt -yq install imagemagick
apt -yq install golang

mkdir -p /opt
cd /opt
git clone https://github.com/transcriptaze/png2ascii
cd png2ascii
make build
export PATH=$PATH:/opt/png2ascii/bin

mkdir -p /opt/mp42ascii
mkdir -p /opt/mp42ascii/frames
mkdir -p /opt/mp42ascii/out

# NOTE: edit memory/size policy in /etc/ImageMagick-6/policy.xml:
# 
# <policy domain="resource" name="memory" value="4GB"/>
#  <policy domain="resource" name="map" value="512MiB"/>
#  <policy domain="resource" name="width" value="32KP"/>
#  <policy domain="resource" name="height" value="32KP"/>
#  <policy domain="resource" name="area" value="1GP"/>
#  <policy domain="resource" name="disk" value="1GiB"/>

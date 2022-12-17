#/bin/bash

# Downloads an MP4 from Google Drive and saves it to /opt/mp42ascii/movie.mov

FILEID="$1"
MP4="movie.mov"

cd /opt/mp42ascii

COOKIES=$(wget --quiet --save-cookies /tmp/cookies.txt --keep-session-cookies --no-check-certificate "https://docs.google.com/uc?export=download&id=${FILEID}" -O- | sed -rn 's/.*confirm=([0-9A-Za-z_]+).*/\1\n/p')
wget --load-cookies /tmp/cookies.txt "https://docs.google.com/uc?export=download&confirm=${COOKIES}&id=${FILEID}" -O "${MP4}" && rm -rf /tmp/cookies.txt

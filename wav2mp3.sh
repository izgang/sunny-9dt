#!/bin/bash
# https://www.google.com/search?q=linux+convert+wav+to+mp3
# http://stackoverflow.com/questions/11216445/converting-wav-files-to-high-quality-mp3-using-linux-terminal
# http://superuser.com/questions/507386/why-would-i-choose-libav-over-ffmpeg-or-is-there-even-a-difference
# http://askubuntu.com/questions/432542/is-ffmpeg-missing-from-the-official-repositories-in-14-04
# sudo apt-get install ffmpeg
# run this script with the following command on Ubuntu 15.10
# $ bash wav2mp3.sh
# the following command is not working:
# $ sh wav2mp3.sh

OUTPUTDIR="output"

mkdir -p $OUTPUTDIR
for i in {000..094}
do
  ffmpeg -i 9dt/9dt$i.wav $OUTPUTDIR/9dt$i.mp3
done

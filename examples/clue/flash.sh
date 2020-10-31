set -e
if [ ! -f '/Volumes/FTHR840BOOT/INDEX.HTM' ]; then
    echo "Clue not mounted. Mounting..."
    [ -d  /Volumes/FTHR840BOOT ] || sudo mkdir /Volumes/FTHR840BOOT
    sudo mount -t msdos /dev/disk2 /Volumes/FTHR840BOOT
fi
tinygo flash -target=clue_alpha .
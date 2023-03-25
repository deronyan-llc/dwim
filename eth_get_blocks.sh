# get ethereum blocks
ROOT=/media/alan/data/columbo-data/blocks/ethereum/mainnet
for n in `seq 1660000 16700000`; do
    h=`printf "%x" $n`;
    FILE=$ROOT/block$n.json
    if [ -f "$FILE" ]; then
        echo -n "."
    else
        echo "Getting $FILE..."
        curl https://distinguished-sly-liquid.discover.quiknode.pro/7684eee5885d6f6e7b8e3f88c60004facf139208/ \
         -X POST \
         --silent \
         -H "Content-Type: application/json" \
         --data "{\"method\":\"eth_getBlockByNumber\",\"params\":[\"0x$h\",true],\"id\":$n,\"jsonrpc\":\"2.0\"}" \
         -o $FILE
	if [ $? -ne 0 ]; then
	    echo "$n" >> ./failed_blocks.out
	fi
	sleep 2;
    fi
done

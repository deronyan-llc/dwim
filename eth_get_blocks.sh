# get ethereum blocks
#set -x
ROOT=/mnt/data/columbo-data/blocks/ethereum/mainnet
count=0
dup=0
ok=0
fail=0
total=0
for n in `seq 16670680 16700000`; do
    h=`printf "%x" $n`;
    FILE=$ROOT/block$n.json
    #if [ -f "$FILE" ]; then
    #    echo -n "_"
	#    dup=$((dup+1))
    #else
        curl https://distinguished-sly-liquid.discover.quiknode.pro/7684eee5885d6f6e7b8e3f88c60004facf139208/ \
             -X POST \
             --silent \
             -H "Content-Type: application/json" \
             --data "{\"method\":\"eth_getBlockByNumber\",\"params\":[\"0x$h\",true],\"id\":$n,\"jsonrpc\":\"2.0\"}" \
             -o $FILE
	    if [ $? -ne 0 ]; then
    	    echo "$n" >> ./failed_blocks.out
	        echo -n "x"
    	    fail=$((fail+1))
	    else
	        echo -n "+"
    	    ok=$((ok+1))
	    fi
    #fi
    count=$((count+1))
    total=$((total+1))
    mod=$((total%3))
    if [ $count -eq 0 ]; then
        sleep 1
    fi
    if [ $count -eq 50 ]; then
	    echo "  $dup Dup, $ok Ok, $fail Fail, $total Total"
	    count=0
    fi
done

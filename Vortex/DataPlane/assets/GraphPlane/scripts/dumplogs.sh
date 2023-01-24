#/bin/bash

echo "Dumping GraphPlane Logs.."

for entry in "../cmd/GraphPlane/logs"/*;do
    if [ "${entry: -4}" == ".log" ]; then
        > $entry
    fi
done
echo "Finished."
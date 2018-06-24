#!/usr/bin/env bash
for d in ${PROJECT_DIR}/*/ ; do
    cd $d

    if [ -f Gopkg.toml ]
    then
    	dep ensure
    else
      dep init
    fi

done

cd ${PROJECT_DIR}

exec "$@"

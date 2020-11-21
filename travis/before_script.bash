#!/usr/bin/env bash

clickhouse-client -q "CREATE TABLE default.internal_logs ( date Date DEFAULT toDate(time), time DateTime, server String, port Int32, type String, table String, volume Int64, message String, content String, source String, server_time DateTime DEFAULT now() ) ENGINE = MergeTree() PARTITION BY tuple() ORDER BY time SETTINGS index_granularity = 8192;"
clickhouse-client -q "CREATE TABLE default.internal_logs_buffer AS default.internal_logs ENGINE = Buffer(default, internal_logs, 2, 10, 10, 10000000, 10000000, 100000000, 100000000);"
clickhouse-client -q "CREATE TABLE default.daemon_heartbeat ( date Date DEFAULT toDate(time), time DateTime, daemon String, server String, port Int32, build_version String, build_commit String, config_ts DateTime, config_hash String, memory_used_bytes Int64, cpu_user_avg Float32, cpu_sys_avg Float32, server_time DateTime DEFAULT now() ) ENGINE = ReplacingMergeTree(date, (daemon, server, port), 8192, server_time);"
clickhouse-client -q "CREATE TABLE default.daemon_heartbeat_buffer AS default.daemon_heartbeat ENGINE = Buffer(default, daemon_heartbeat, 2, 15, 15, 10000000, 10000000, 100000000, 100000000);"

go get github.com/VKCOM/kittenhouse
cd $GOPATH/bin/kittenhouse
./kittenhouse -u="" -g=""
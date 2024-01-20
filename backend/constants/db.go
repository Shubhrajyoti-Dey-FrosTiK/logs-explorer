package constants

import "time"

const DB = "test"
const COLLECTION_LOGS = "logs"
const CONNECTION_STRING = "ATLAS_URI"

const CACHING_DURATION = 20 * time.Minute
const CACHE_CONTROL_HEADER = "Cache-Control"
const NO_CACHE = "no-cache"

const REDIS_LOG_QUEUE = "REDIS_LOG_QUEUE"
const REDIS_LOG_STATUS = "REDIS_LOG_STATUS"

var FIND_LIMIT int = 60

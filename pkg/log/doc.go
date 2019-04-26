/*
Package log is the kratos log library.

First, the main function:

1. The log is printed to the local
2. Log print to standard output
3. verbose log implementation. Refer to the glog implementation. You can set different verbose levels. The default is not enabled.

Second, the log configuration

Default configuration

The current log has implemented the default configuration. You can use the following methods directly:
log.Init(nil)

2. Start parameter or environment variable

Startup parameters Environment variable Description
Log.stdout LOG_STDOUT Whether to enable standard output
Log.dir LOG_DIR file log path
Log.v LOG_V verbose log level
Log.module LOG_MODULE The verbose level of each file can be configured separately: file=1, file2=2
Log.filter LOG_FILTER Configure the fields to be filtered: field1,field2

3. Configuration file
But if you have special needs, you can take a look at the format configuration:
[log]
Family = "xxx-service"
Dir = "/data/log/xxx-service/"
Stdout = true
vLevel = 3
Filter = ["fileld1", "field2"]
[log.module]
"dao_user" = 2
"servic*" = 1

Third, the configuration instructions

1.log

Family project name, default read environment variable $APPID
Studout standard output, prod environment is not recommended to open
Filter configures the fields that need to be filtered out and replaces them with "***"
Dir file log address, prod environment is not recommended to open
v Enable the verbose level log to specify the global level.

2. log.module

The verbose level of each file can be configured separately
*/
package log

redis-cli.o: redis-cli.c fmacros.h version.h ../deps/hiredis/hiredis.h \
  ../deps/hiredis/read.h ../deps/hiredis/sds.h ../deps/hiredis/alloc.h \
  ../deps/hiredis/sdscompat.h dict.h adlist.h zmalloc.h \
  ../deps/linenoise/linenoise.h help.h anet.h ae.h monotonic.h \
  cli_common.h

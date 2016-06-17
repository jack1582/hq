# -ldflags "-s" 忽略debug的打印信息
#  -gcflags "-N -l" 忽略go编译内部优化，避免为gdb调试带来困难
#  用make，只是个习惯

GXX=go build -gcflags "-N -l"

hq: hq.go
	$(GXX) $^

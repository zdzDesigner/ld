

test:
	@bash ./tests/init.sh
	@printf "\e[32mOK\e[0m\n"



## optimize
ld:
	@go build -ldflags="-s -w" -o $@ && upx -9 $@

elf:
	@hexdump -C out/init/out.o

# make elfn BYTE=100
elfn:
	@hexdump -C -n $(BYTE) out/init/out.o


section:
	@readelf -S  out/init/out.o

.PHONY: test elfn ld

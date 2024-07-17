


drivers:=default sqlite3 mysql postgres

defaultSQLFiles      :=$(wildcard *.default.sql)
sqlite3SQLFiles      :=$(wildcard *.sqlite3.sql)
mysqlSQLFiles        :=$(wildcard *.mysql.sql)
postgresSQLFiles   :=$(wildcard *.postgres.sql)

defaultEmbedFiles    :=$(patsubst %,%.go,$(defaultSQLFiles))
sqlite3EmbedFiles    :=$(patsubst %,%.go,$(sqlite3SQLFiles))
mysqlEmbedFiles      :=$(patsubst %,%.go,$(mysqlSQLFiles))
postgresEmbedFiles :=$(patsubst %,%.go,$(postgresSQLFiles))

MapFiles             :=$(patsubst %,%.map.go,$(drivers))
EmbedFiles           :=$(defaultEmbedFiles) $(defaultEmbedFiles) $(mysqlEmbedFiles) $(postgresEmbedFiles)
SQLFiles             :=$(defaultSQLFiles) $(defaultSQLFiles) $(mysqlSQLFiles) $(postgresSQLFiles)


all: build

build: $(MapFiles)

#%.map.go: $(GOFiles)
%.map.go: $($(patsubst %.map.go,%,$@)EmbedFiles) Makefile
	@echo map $(patsubst %.map.go,%,$@)
	@ls *.$(patsubst %.map.go,%,$@).sql 2>/dev/null|sed -E 's/.[^.]+.sql$$//g'|xargs -I {} ls {}.default.sql >/dev/null 2>&1 ; if [ $$? -ne 0 ];then echo Engined SQL file without default;false;fi
	@if (ls *.sql 2>/dev/null|sed -E 's/^[0-9]\{4\}\.//g'|sed -E 's/.sql$$//g'|grep -v default|grep -v sqlite3|grep -v postgres|grep -v -q mysql ) ; then echo UNExpect SQL FILE ;false; fi
	@rm -fr $@
	@ls *.$(patsubst %.map.go,%,$@).sql 2>/dev/null|awk 'BEGIN{E="$(patsubst %.map.go,%,$@)"; print("package QuerySQL\nfunc load_"E"_map(){")}  {F=$$1;K=F;V=F; gsub(/\.[^.]+\.[^.]+$$/,"",K); gsub("\\.","_",V); print( E"_map[\""K"\"]=str_"V);} END{print "}"}' >> $@

%.sql.go: %.sql Makefile
	@echo embed $<
	@rm -fr $@
	@echo "package QuerySQL" >>$@
	@echo "import _ \"embed\"" >>$@
	@echo "" >>$@
	@echo "//go:embed $<" >>$@
	@echo "var str_$(subst .,_,$<) string" >>$@
	@echo "" >>$@



clean:
	rm -fr *.sql.go *.map.go


.PRECIOUS: %.sql.go

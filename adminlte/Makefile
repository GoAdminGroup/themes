CLI=adm
ASSETS_PATH=./resource/assets

all: assets

assets:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	rm -rf $(ASSETS_PATH)/dist
	mkdir $(ASSETS_PATH)/dist
	mkdir $(ASSETS_PATH)/dist/js
	mkdir $(ASSETS_PATH)/dist/css
	cp $(ASSETS_PATH)/src/js/*.js $(ASSETS_PATH)/dist/js/
	cp $(ASSETS_PATH)/src/css/*.png $(ASSETS_PATH)/dist/css/
	cp -R $(ASSETS_PATH)/src/css/fonts $(ASSETS_PATH)/dist/css/
	cp -R $(ASSETS_PATH)/src/img $(ASSETS_PATH)/dist/
	cp -R $(ASSETS_PATH)/src/fonts $(ASSETS_PATH)/dist/
	make combine
	$(CLI) compile asset
	make build-tmpl

combine:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	make combine-js
	make combine-css

combine-js:
	$(CLI) combine js --src=$(ASSETS_PATH)/src/js/all/ --dist=$(ASSETS_PATH)/dist/js/all.min.js
	$(CLI) combine js --src=$(ASSETS_PATH)/src/js/all_2/ --dist=$(ASSETS_PATH)/dist/js/all_2.min.js
	$(CLI) combine js --src=$(ASSETS_PATH)/src/js/form/ --dist=$(ASSETS_PATH)/dist/js/form.min.js
	$(CLI) combine js --src=$(ASSETS_PATH)/src/js/content/ --dist=$(ASSETS_PATH)/dist/js/content.min.js

combine-css:
	$(CLI) combine css

build-tmpl:
	$(CLI) compile tpl -p=adminlte
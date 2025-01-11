build-web:
	cd ./web; \
		NODE_ENV="production" pnpm generate; \
		cp -r .output/public/. ../api/dist

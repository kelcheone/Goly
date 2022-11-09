## Get started

### Spinnup the postgres Db

```bash
cd app
docker compose up -d
```

### Install the dependencies...

> Go
>
> ```bash
> go mod tidy
> ```
>
> Svelte
>
> ```bash
> cd client
> yarn install
> ```

...then start [Rollup](https://rollupjs.org):

```bash
yarn dev
```

import App from './App.svelte';
import { SpaRouter } from 'svelte-router-spa'
import { routes } from './routes'
import NotFound from './views/not_found.svelte'

SpaRouter({
  routes,
  pathName: document.location.href,
  notFound: NotFound
}).getActiveRoute

const app = new App({
	target: document.body,
	props: {
		name: 'world'
	}
});

export default app;
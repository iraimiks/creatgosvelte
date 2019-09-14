import Home from './views/Home.svelte'
import About from './views/About.svelte'
const routes = [
  {
    name: '/',
    component: Home
  },
  {
    name: 'About',
    component: About
  }
]

export { routes }
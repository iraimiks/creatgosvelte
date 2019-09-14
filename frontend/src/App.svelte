<div id="app">
  <Header></Header>
  <div id="main">
    <Router />
  </div>
  <Footer></Footer>
</div>

<script>
	export let name;
    import { Router } from 'svelte-router-spa';
    import Header from './templates/Header.svelte'
    import Footer from './templates/Footer.svelte'
</script>

<style lang="scss">
 :global(*) {
   margin: 0;
   padding: 0;
   border: 0;
   outline: none;
 }
 /*Shto eta*/
 :global(html) {
   height: 0;
 }

 :global(body) {
   min-height: 100%;
 }
  #main {
    padding: 0;
  }

  #app {
  min-height: 100%;
      display: grid;
      grid-template-areas:
        "header"
        "main"
        "footer";
  }
  #main {
    grid-area: main;
  }
</style>

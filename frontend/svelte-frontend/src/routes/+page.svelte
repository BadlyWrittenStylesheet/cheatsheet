<script>
  import CheatsheetMenu from '$lib/components/CheatsheetMenu.svelte';
  
  export let data;
  let selectedContent = '';
  
  async function loadCheatsheet(id) {
    const response = await fetch(`http://localhost:55003/cheatsheets/${id}`);
    const content = await response.text();
    selectedContent = content;
  }
</script>

<div class="layout">
  <nav>
    <CheatsheetMenu 
      cheatsheets={data.cheatsheets} 
      onSelect={loadCheatsheet}
    />
  </nav>
  
  <main class="content">
    {#if selectedContent}
        {@html selectedContent}
    {:else}
      <p>Select a cheatsheet from the nav</p>
    {/if}
  </main>
</div>

<style>
  .layout {
    display: grid;
    grid-template-columns: auto 1fr;
    height: 100vh;
  }

  nav {
    border-right: 1px solid #eee4;
    height: 100%;
    overflow-y: auto;
  }

  main {
    padding: 0 2rem;
    overflow-y: auto;
  }

  main :global(p) {
    white-space: pre-line; /* xd */
  }
</style>

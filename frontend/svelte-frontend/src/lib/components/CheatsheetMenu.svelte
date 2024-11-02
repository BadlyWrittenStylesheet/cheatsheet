<script>
    export let cheatsheets = [];
    export let onSelect = (id) => {};
</script>

<div class="menu">
  {#each cheatsheets as cheatsheet}
    <button 
      class="menu-item" 
      on:click={() => onSelect(cheatsheet.id)}
    >
      {cheatsheet.title}
    </button>
  {/each}
</div>

<style>
  .menu {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
    padding-right: 2rem;
  }

  .menu-item {
    text-align: left;
    padding: 0.5rem .75rem;
    border: none;
    background: #e4e4e4;
    border-radius: 6px;
    transition: 300ms;
  }

  .menu-item:hover {
    background: #ff8a8a;
  }
</style>

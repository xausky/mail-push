<script>
  export let tabs = [];
  export let activeTab = '';
  
  function switchTab(tab) {
    activeTab = tab;
  }
</script>

<div class="tabs">
  {#each tabs as tab}
    <button 
      class:active={activeTab === tab.id} 
      on:click={() => switchTab(tab.id)}
    >
      {tab.label}
    </button>
  {/each}
</div>

<div class="tab-content">
  <slot></slot>
</div>

<style>
  .tabs {
    display: flex;
    margin-bottom: 15px;
    border-bottom: 1px solid #ddd;
  }
  
  .tabs button {
    background-color: transparent;
    color: #333;
    padding: 10px 20px;
    border: none;
    border-bottom: 2px solid transparent;
    cursor: pointer;
    transition: all 0.3s ease;
  }
  
  .tabs button:hover {
    background-color: #f0f0f0;
  }
  
  .tabs button.active {
    color: #4CAF50;
    border-bottom: 2px solid #4CAF50;
    background-color: transparent;
  }
  
  .tab-content {
    padding: 15px 0;
  }
</style> 
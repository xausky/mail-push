<script>
  export let id = '';
  export let label = '';
  export let value = '';
  export let options = [];
  export let placeholder = '';
</script>

<div class="form-group">
  <label for={id}>{label}</label>
  <select 
    {id} 
    bind:value 
    class:placeholder={!value}
  >
    {#each options as option}
      <option value={option.key}>{option.name}</option>
    {/each}
  </select>
</div>

<style>
  .form-group {
    margin-bottom: 15px;
    display: flex;
    align-items: center;
    gap: 10px;
  }
  
  label {
    min-width: 100px;
    text-align: right;
    color: #333;
    font-size: 14px;
  }
  
  select {
    flex: 1;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 14px;
    transition: border-color 0.3s ease;
    background-color: white;
    cursor: pointer;
  }
  
  select:focus {
    outline: none;
    border-color: #4CAF50;
    box-shadow: 0 0 0 2px rgba(76, 175, 80, 0.1);
  }
  
  select.placeholder {
    color: #999;
  }
  
  option {
    color: #333;
  }
</style>
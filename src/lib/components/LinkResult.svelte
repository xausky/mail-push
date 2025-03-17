<script>
  export let link = '';
  export let onCopy;
  export let copyButtonText = '复制';
</script>

<div class="result-container">
  <p>{link}</p>
  <button on:click={onCopy}>{copyButtonText}</button>
</div>

<style>
  .result-container {
    display: flex;
    align-items: center;
    gap: 10px;
    margin: 15px 0;
  }
  
  .result-container p {
    word-break: break-all;
    margin: 0;
    flex-grow: 1;
  }
  
  button {
    background-color: #4CAF50;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 14px;
    transition: background-color 0.3s ease;
  }
  
  button:hover {
    background-color: #45a049;
  }
</style> 
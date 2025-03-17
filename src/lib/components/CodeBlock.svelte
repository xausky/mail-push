<script>
  export let code = '';
  export let language = 'curl';
  export let onCopy;
  export let copyButtonText = '复制';
</script>

<div class="code-block-container">
  <div class="code-block-header">
    <span>{language}</span>
    <button on:click={onCopy}>{copyButtonText}</button>
  </div>
  <pre class="code-block">{@html code}</pre>
</div>

<style>
  .code-block-container {
    position: relative;
    margin: 15px 0;
    border-radius: 6px;
    overflow: hidden;
    border: 1px solid #ddd;
  }
  
  .code-block-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #f0f0f0;
    padding: 8px 15px;
    border-bottom: 1px solid #ddd;
    font-family: monospace;
    font-size: 12px;
    color: #555;
  }
  
  .code-block-header button {
    padding: 4px 8px;
    font-size: 12px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .code-block-header button:hover {
    background-color: #45a049;
  }
  
  .code-block {
    margin: 0;
    padding: 15px;
    background-color: #fff;
    overflow-x: auto;
    font-family: monospace;
    font-size: 14px;
    line-height: 1.5;
  }
</style> 
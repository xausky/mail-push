<script>
  import LinkResult from './LinkResult.svelte';
  import CodeBlock from './CodeBlock.svelte';
  
  export let format = {};
  export let link = '';
  export let exampleCode = '';
  export let onCopyLink;
  export let onCopyExample;
  export let copyLinkButtonText = '复制';
  export let copyExampleButtonText = '复制';
</script>

<div>
  <h4>{format.name} 兼容链接</h4>
  <LinkResult 
    link={link} 
    onCopy={onCopyLink} 
    copyButtonText={copyLinkButtonText} 
  />
  <p>此地址可直接配置到支持 {format.name} 的 Webhook 地址中使用。</p>
    <p>详细的消息格式说明请参考：<a href={format.document_url} target="_blank">{format.name} 文档</a></p>
  <h5>请求示例：</h5>
  <CodeBlock 
    code={exampleCode} 
    language="curl" 
    onCopy={onCopyExample} 
    copyButtonText={copyExampleButtonText} 
  />
</div>

<style>
  h4 {
    margin-top: 0;
    color: #4CAF50;
  }
  
  h5 {
    margin-bottom: 5px;
  }
</style>
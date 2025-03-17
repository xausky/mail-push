<script>
  import LinkResult from './LinkResult.svelte';
  import CodeBlock from './CodeBlock.svelte';
  
  export let link = '';
  export let exampleCode = '';
  export let onCopyLink;
  export let onCopyExample;
  export let copyLinkButtonText = '复制';
  export let copyExampleButtonText = '复制';
</script>

<div>
  <h4>Discord 兼容链接</h4>
  <LinkResult 
    link={link} 
    onCopy={onCopyLink} 
    copyButtonText={copyLinkButtonText} 
  />
  <p>此地址可直接配置到支持 Discord 的 Webhook 地址中使用。</p>
  <h5>请求示例：</h5>
  <CodeBlock 
    code={exampleCode} 
    language="curl" 
    onCopy={onCopyExample} 
    copyButtonText={copyExampleButtonText} 
  />
  <p>详细的消息格式说明请参考：<a href="https://discord.com/developers/docs/resources/webhook#execute-webhook" target="_blank">Discord Webhook 文档</a></p>
  <p>支持的字段：</p>
  <ul>
    <li>content - 消息内容</li>
    <li>username - 自定义机器人名称</li>
    <li>avatar_url - 自定义机器人头像</li>
    <li>embeds - 嵌入式内容</li>
  </ul>
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
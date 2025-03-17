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
  <h4>企业微信兼容链接</h4>
  <LinkResult 
    link={link} 
    onCopy={onCopyLink} 
    copyButtonText={copyLinkButtonText} 
  />
  <p>此地址可直接配置到支持企业微信群机器人的 Webhook 地址中使用。</p>
  <h5>请求示例：</h5>
  <CodeBlock 
    code={exampleCode} 
    language="curl" 
    onCopy={onCopyExample} 
    copyButtonText={copyExampleButtonText} 
  />
  <p>详细的消息格式说明请参考：<a href="https://developer.work.weixin.qq.com/document/path/91770" target="_blank">企业微信群机器人配置说明</a></p>
  <p>支持的消息类型：</p>
  <ul>
    <li>文本消息 (text)</li>
    <li>markdown消息 (markdown)</li>
    <li>图片消息 (image)</li>
    <li>图文消息 (news)</li>
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
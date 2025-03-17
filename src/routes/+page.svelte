<script lang="ts">
  import FormInput from '$lib/components/FormInput.svelte';
  import TabPanel from '$lib/components/TabPanel.svelte';
  import WechatTab from '$lib/components/WechatTab.svelte';
  import DiscordTab from '$lib/components/DiscordTab.svelte';
  import ProjectInfo from '$lib/components/ProjectInfo.svelte';
  
  // 状态变量
  let provider = '';
  let mailUser = '';
  let mailToken = '';
  let wechatLink = '';
  let discordLink = '';
  let showResult = false;
  let activeTab = 'wechat'; // 默认显示企业微信标签页
  
  // 标签页配置
  const tabs = [
    { id: 'wechat', label: '企业微信' },
    { id: 'discord', label: 'Discord' }
  ];
  
  // 生成链接函数
  function generateLink() {
    if (!provider || !mailUser || !mailToken) {
      alert('请填写所有字段！');
      return;
    }
    
    const data = btoa(`${provider}|${mailUser}|${mailToken}`);
    wechatLink = `${window.location.origin}/send/${data}`;
    discordLink = `${window.location.origin}/discord/${data}`;
    showResult = true;
  }
  
  // 复制到剪贴板函数
  async function copyToClipboard(text: string): Promise<boolean> {
    try {
      await navigator.clipboard.writeText(text);
      return true;
    } catch (err) {
      console.error('复制失败:', err);
      return false;
    }
  }
  
  // 处理复制按钮点击
  let copyButtonText = {
    wechat: '复制',
    discord: '复制',
    wechatExample: '复制',
    discordExample: '复制'
  };
  
  type CopyType = 'wechat' | 'discord' | 'wechatExample' | 'discordExample';
  
  async function handleCopy(type: CopyType) {
    let textToCopy = '';
    let buttonType = type;
    
    if (type === 'wechat') {
      textToCopy = wechatLink;
    } else if (type === 'discord') {
      textToCopy = discordLink;
    } else if (type === 'wechatExample') {
      textToCopy = getWechatCurlExample();
      buttonType = 'wechatExample';
    } else if (type === 'discordExample') {
      textToCopy = getDiscordCurlExample();
      buttonType = 'discordExample';
    }
    
    const success = await copyToClipboard(textToCopy);
    
    if (success) {
      copyButtonText[buttonType] = '已复制！';
      setTimeout(() => {
        copyButtonText[buttonType] = '复制';
      }, 2000);
    } else {
      copyButtonText[buttonType] = '复制失败';
      setTimeout(() => {
        copyButtonText[buttonType] = '复制';
      }, 2000);
    }
  }

  // 获取当前示例代码
  function getWechatCurlExample() {
    return `# 文本消息 curl 示例
curl -X POST "${wechatLink}" \\
     -H "Content-Type: application/json" \\
     -d '{
        "msgtype": "text",
        "text": {
            "content": "这是一条文本消息"
        }
}'`;
  }

  function getDiscordCurlExample() {
    return `# Discord Webhook 示例
curl -X POST "${discordLink}" \\
     -H "Content-Type: application/json" \\
     -d '{
        "content": "这是一条 Discord 消息",
        "username": "Mail Push Bot",
        "avatar_url": "https://i.imgur.com/4M34hi2.png"
     }'`;
  }
</script>

<main>
  <h1>Mail Push</h1>
  
  <FormInput 
    id="provider" 
    label="邮件服务商:" 
    bind:value={provider} 
    placeholder="例如: 163,qq" 
  />
  
  <FormInput 
    id="mail_user" 
    label="用户名:" 
    bind:value={mailUser} 
    placeholder="输入邮箱用户名（不含@后缀）" 
    autocomplete="off" 
  />
  
  <FormInput 
    id="mail_token" 
    label="密码或授权码:" 
    bind:value={mailToken} 
    placeholder="输入邮箱密码或授权码（明文显示）" 
    autocomplete="off" 
  />
  
  <div class="button-container">
    <button on:click={generateLink}>生成链接</button>
  </div>
  
  {#if showResult}
    <div id="result">
      <TabPanel tabs={tabs} bind:activeTab>
        {#if activeTab === 'wechat'}
          <WechatTab 
            link={wechatLink}
            exampleCode={getWechatCurlExample()}
            onCopyLink={() => handleCopy('wechat')}
            onCopyExample={() => handleCopy('wechatExample')}
            copyLinkButtonText={copyButtonText.wechat}
            copyExampleButtonText={copyButtonText.wechatExample}
          />
        {:else if activeTab === 'discord'}
          <DiscordTab 
            link={discordLink}
            exampleCode={getDiscordCurlExample()}
            onCopyLink={() => handleCopy('discord')}
            onCopyExample={() => handleCopy('discordExample')}
            copyLinkButtonText={copyButtonText.discord}
            copyExampleButtonText={copyButtonText.discordExample}
          />
        {/if}
      </TabPanel>
    </div>
  {/if}
  
  <ProjectInfo />
</main>

<style>
  main {
    font-family: Arial, sans-serif;
    max-width: 800px;
    margin: 20px auto;
    padding: 0 20px;
  }
  
  .button-container {
    margin-top: 20px;
    text-align: center;
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
  
  #result {
    margin-top: 20px;
    padding: 15px;
    border: 1px solid #ddd;
    border-radius: 4px;
    background-color: #f8f9fa;
  }
</style>

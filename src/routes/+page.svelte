<script lang="ts">
  import FormInput from '$lib/components/FormInput.svelte';
  import FormSelect from '$lib/components/FormSelect.svelte';
  import TabPanel from '$lib/components/TabPanel.svelte';
  import WebhookTab from '$lib/components/WebhookTab.svelte';
  import ProjectInfo from '$lib/components/ProjectInfo.svelte';
  import {fetchMailProviders, fetchWebhookFormats, type MailProvider, type WebhookFormat} from '$lib/api';
  import { onMount } from 'svelte';
  
  // 状态变量
  let provider = '';
  let mailUser = '';
  let mailToken = '';
  let showResult = false;
  let activeTab = 'wechat'; // 默认显示企业微信标签页
  
  // 动态数据
  let mailProviders: MailProvider[] = [];
  let webhookFormats: WebhookFormat[] = [];
  let tabs: any[] = [];
  
  // 页面加载时获取数据
  onMount(async () => {
    // 获取邮件服务商列表
    mailProviders = await fetchMailProviders();
    provider = mailProviders[0].key
    
    // 获取WebHook格式列表
    webhookFormats = await fetchWebhookFormats();
    
    // 动态生成标签页
    tabs = webhookFormats.map(format => ({
      id: format.key,
      label: format.name
    }));
    
    // 设置默认激活的标签页
    if (tabs.length > 0) {
      activeTab = tabs[0].id;
    }
  });
  
  // 生成链接函数
  function generateLink() {
    if (!provider || !mailUser || !mailToken) {
      alert('请填写所有字段！');
      return;
    }

    // 初始化复制按钮文本
    initCopyButtonText();
    
    showResult = true;
  }
  
  // 获取WebHook链接
  function getWebhookLink(key: string) {
    const data = urlSafeBase64Encode(`${provider}|${mailUser}|${mailToken}`);
    return `${window.location.origin}/${key}/${data}`;
  }

  function urlSafeBase64Encode(str: string) {
    // 1. 先进行标准Base64编码
    const base64 = btoa(str);

    // 2. 替换URL不安全的字符
    const urlSafe = base64
      .replace(/\+/g, '-')  // 将 + 替换为 -
      .replace(/\//g, '_'); // 将 / 替换为 _

    // 3. 移除填充字符 =
    return urlSafe.replace(/=/g, '');
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
  let copyButtonText: Record<string, string> = {};
  
  // 初始化复制按钮文本
  function initCopyButtonText() {
    webhookFormats.forEach(format => {
      copyButtonText[format.key] = '复制';
      copyButtonText[`${format.key}Example`] = '复制';
    });
  }
  
  async function handleCopy(type: string) {
    let textToCopy = '';
    let buttonType = type;

    if (type.endsWith('Example')) {
      const formatKey = type.replace('Example', '');
      textToCopy = getWebhookExample(formatKey);
    } else {
      textToCopy = getWebhookLink(type);
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
  function getWebhookExample(key: string) {
    const format = webhookFormats.find(f => f.key === key);
    if (!format) return '';
    
    // 替换示例中的占位符
    let example = format.example;
    const baseUrl = window.location.origin;
    const data = urlSafeBase64Encode(`${provider}|${mailUser}|${mailToken}`);
    
    example = example.replace('<BASE_URL>', baseUrl);
    example = example.replace('<BASE64_DATA>', data);
    
    return example;
  }
</script>

<main>
  <h1>Mail Push</h1>
  
  <FormSelect
    id="provider"
    label="邮件服务商:"
    bind:value={provider}
    options={mailProviders}
    placeholder="请选择邮件服务商"
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
        {#each webhookFormats as format}
          {#if activeTab === format.key}
              <WebhookTab
                      format={format}
                      link={getWebhookLink(format.key)}
                      exampleCode={getWebhookExample(format.key)}
                      onCopyLink={() => handleCopy(format.key)}
                      onCopyExample={() => handleCopy(`${format.key}Example`)}
                      copyLinkButtonText={copyButtonText[format.key] || '复制'}
                      copyExampleButtonText={copyButtonText[`${format.key}Example`] || '复制'}
              />
          {/if}
        {/each}
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

// API服务，用于获取邮件服务商和WebHook格式数据

// 邮件服务商类型
export interface MailProvider {
  key: string;
  name: string;
}

// WebHook格式类型
export interface WebhookFormat {
  key: string;
  name: string;
  document_url: string;
  example: string;
}

// 获取邮件服务商列表
export async function fetchMailProviders(): Promise<MailProvider[]> {
  try {
    const response = await fetch('/providers');
    if (!response.ok) {
      throw new Error(`获取邮件服务商失败: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error('获取邮件服务商出错:', error);
    return [];
  }
}

// 获取WebHook格式列表
export async function fetchWebhookFormats(): Promise<WebhookFormat[]> {
  try {
    const response = await fetch('/formats');
    if (!response.ok) {
      throw new Error(`获取WebHook格式失败: ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error('获取WebHook格式出错:', error);
    return [];
  }
}
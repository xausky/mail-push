<div class="docs">
  <h2>项目简介</h2>
  <p>Mail Push 是一个基于 SMTP 协议的消息推送服务，<strong>只能</strong>发送邮件到配置的邮箱地址（即你自己的邮箱）来实现消息推送功能。这是一个单向的推送服务，不支持发送到其他邮箱地址。</p>
  
  <h3>工作原理</h3>
  <ul>
    <li>使用 SMTP 协议发送邮件到<strong>配置的邮箱地址</strong></li>
    <li>邮件的发件人和收件人是<strong>同一个邮箱</strong></li>
    <li>可通过各种邮件客户端接收提醒</li>
    <li>特别地，使用 QQ 邮箱可以获得更好的体验：
      <ul>
        <li>启用微信的 QQ 邮箱提醒功能，获得更实时的通知</li>
        <li>直接在微信中查看邮件内容</li>
        <li>开启方式：微信->我->设置->通用->辅助功能->QQ邮箱提醒</li>
      </ul>
    </li>
  </ul>

  <h2>安全说明</h2>
  <ul>
    <li>所有认证信息使用 Base64 编码存储在 URL 中</li>
    <li>服务器不存储任何用户信息</li>
    <li>建议自行部署服务以确保信息安全</li>
  </ul>
  
  <h2>开源地址</h2>
  <p>本项目已在 GitHub 开源：<a href="https://github.com/xausky/mail-push" target="_blank">github.com/xausky/mail-push</a> 因为有可能被墙被D被删，所以有高可用或者高隐私需求的可以自行搭建。</p>
</div>

<style>
  .docs {
    margin-top: 30px;
    padding: 15px;
    background-color: #f8f9fa;
    border-radius: 4px;
  }
</style> 
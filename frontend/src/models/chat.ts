export class ChatMessage {
  constructor(
    public content = '',
    public createdAt = new Date().toISOString(),
    public from = '',
    public to = ''
  ) {}
}

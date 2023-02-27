export class Chat {
  constructor(
    public content: string,
    public createdAt = new Date().toISOString(),
    public from = '',
    public to = ''
  ) {}
}

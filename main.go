package main

func main() {
	client, ctx, cancel := initMongo()
	api := initAPI(ctx, client)

	defer cancel()
	defer client.Disconnect(ctx)

	api.Run()
}

# This will use dposc

# FIXME: need to import dp_osc and osc_request corectly

def main():
    client = OSCStreamingClient()
    client.connect((osc_host, osc_port))

    dp = DigitalPerformer(client)

    while True:
        # do stuff here
        pass

    client.close()


if __name__ == '__main__':
    main()

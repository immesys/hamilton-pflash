# hamilton-pflash

The production-flash (pflash) tool replaces "make flash" for programming a
hamilton mote for deployment.

- Allocates the mote a serial number
- Allocates the mote a MAC address
- Burns an ed25519 keypair, from an newly creates bosswave entity that is in .pflash/<serial>.ent
- Burns a symmetric key
- Burns the mote type (eg 0x7 for hamilton-7)
- Burns the time the mote was flashed
- Locks the mote so that it cannot be debugged.
- Records the github repository and commit of the app

To do this, the following variables must be set:

```
$MOTETYPE e.g. 7 or 3c
$SCRIPTPW (my database password)
optionally $NOSECURE if you want to omit locking the mote
```

The following symbols can be declared in the app to make use of this information

```
const uint64_t* const fb_sentinel     = ((const uint64_t* const)0x3fc00);
const uint64_t* const fb_flashed_time = ((const uint64_t* const)0x3fc08);
const uint8_t*  const fb_mac          = ((const uint8_t*  const)0x3fc10);
const uint16_t* const fb_device_id    = ((const uint16_t* const)0x3fc18);
const uint64_t* const fb_designator   = ((const uint64_t* const)0x3fc1c);
const uint8_t*  const fb_aes128_key   = ((const uint8_t*  const)0x3fc30);
const uint8_t*  const fb_25519_pub    = ((const uint8_t*  const)0x3fc40);
const uint8_t*  const fb_25519_priv   = ((const uint8_t*  const)0x3fc60);
#define FB_SENTINEL_VALUE 0x27c83f60f6b6e7c8
#define HAS_FACTORY_BLOCK (*fb_sentinel == FB_SENTINEL_VALUE)
```

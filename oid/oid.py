import zlib
z = "4acac94f523063c848cdc9c9e702000000ffff"
zlib_deflate = bytes.fromhex(z)
zlib_data = bytes.fromhex("789c4acac94f523063c848cdc9c9e702000000ffff")
deflate_obj = zlib.decompressobj(-zlib.MAX_WBITS)
data_obj = zlib.decompressobj()
print(bytes.decode(deflate_obj.decompress(zlib_deflate), 'utf-8'))
print(bytes.decode(data_obj.decompress(zlib_data), 'utf-8'))

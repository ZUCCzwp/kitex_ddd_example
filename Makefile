generate/user:
	kitex -module github.com/ZUCCzwp/kitex_ddd_example -type protobuf -service userService -I proto proto/user.proto

generate/hello:
	kitex -module github.com/ZUCCzwp/kitex_ddd_example -type protobuf -service helloService -I proto proto/hello.proto
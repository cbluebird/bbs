update_user_kitex_gen:
	kitex -module bbs idl/user.thrift

update_user_rpc:
	kitex -module bbs -service user -use bbs/kitex_gen ../../../idl/user.thrift
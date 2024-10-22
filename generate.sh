protoc  --go_out=./  \
        --go-grpc_out=./ \
        society.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./society.proto

#protoc  --go_out=./  \
#        --go-grpc_out=./ \
        #new_user.proto
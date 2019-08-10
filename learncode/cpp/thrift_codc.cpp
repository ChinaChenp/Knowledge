#include <stdio.h>

#include <protocol/TBinaryProtocol.h>
#include <transport/TTransportUtils.h>
#include <shared_ptr>

using namespace apache::thrift::transport;
using namespace apache::thrift::protocol;

template<typename ThriftStruct>
std::string ThriftToString(const ThriftStruct& ts) {
    std::shared_ptr<TMemoryBuffer> buffer(new TMemoryBuffer);
    boost::shared_ptr<TTransport> trans(buffer);

    TBinaryProtocol protocol(trans);
    ts.write(&protocol);

    uint8_t* buf = NULL;
    uint32_t size = 0;
    buffer->getBuffer(&buf, &size);
    return std::string((char*)buf, (unsigned int)size);
}

template<typename ThriftStruct>
bool StringToThrift(const std::string& buff, ThriftStruct* ts) {
    std::shared_ptr<TMemoryBuffer> buffer(new TMemoryBuffer);
    buffer->write((const uint8_t*)buff.data(), buff.size());

    std::shared_ptr<TTransport> trans(buffer);
    TBinaryProtocol protocol(trans);
    ts->read(&protocol);
    return true;
}

int main() {
    
}
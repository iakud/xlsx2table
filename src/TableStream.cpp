#include "TableStream.h"

#include <cstring>
#include <endian.h>

#if defined(__GNUC__) || defined(__GNUG__)
#	pragma GCC diagnostic ignored "-Wold-style-cast"
#endif

TableStream::TableStream(char* p)
	: _p(p) {
}

TableStream& TableStream::operator>>(bool& b)  {
	b = *reinterpret_cast<bool*>(_p);
	_p += sizeof(bool);
	return *this;
}

TableStream& TableStream::operator>>(int8_t& i) {
	*(int8_t*)&i = *reinterpret_cast<uint8_t*>(_p);
	_p += sizeof(uint8_t);
	return *this;
}

TableStream& TableStream::operator>>(int16_t& i) {
	*(int16_t*)&i = be16toh(*reinterpret_cast<uint16_t*>(_p));
	_p += sizeof(uint16_t);
	return *this;
}

TableStream& TableStream::operator>>(int32_t & i) {
	*(int32_t*)&i = be32toh(*reinterpret_cast<uint32_t*>(_p));
	_p += sizeof(uint32_t);
	return *this;
}

TableStream& TableStream::operator>>(int64_t& i) {
	*(int64_t*)&i = be64toh(*reinterpret_cast<uint64_t*>(_p));
	_p += sizeof(uint64_t);
	return *this;
}

TableStream& TableStream::operator>>(uint8_t& n) {
	n = *reinterpret_cast<uint8_t*>(_p);
	_p += sizeof(uint8_t);
	return *this;
}

TableStream& TableStream::operator>>(uint16_t& n) {
	n = be16toh(*reinterpret_cast<uint16_t*>(_p));
	_p += sizeof(uint16_t);
	return *this;
}

TableStream& TableStream::operator>>(uint32_t& n) {
	n = be32toh(*reinterpret_cast<uint32_t*>(_p));
	_p += sizeof(uint32_t);
	return *this;
}

TableStream& TableStream::operator>>(uint64_t& n) {
	n = be64toh(*reinterpret_cast<uint64_t*>(_p));
	_p += sizeof(uint64_t);
	return *this;
}

TableStream& TableStream::operator>>(float& f) {
	*(uint32_t*)&f = be32toh(*reinterpret_cast<uint32_t*>(_p));
	_p += sizeof(uint32_t);
	return *this;
}

TableStream& TableStream::operator>>(double& f) {
	*(uint64_t*)&f = be64toh(*reinterpret_cast<uint64_t*>(_p));
	_p += sizeof(uint64_t);
	return *this;
}

TableStream& TableStream::operator>>(std::string& s) {
	uint32_t size = be32toh(*reinterpret_cast<uint32_t*>(_p));
	_p += sizeof(uint32_t);
	s.assign(_p, size);
	_p += size;
	return *this;
}
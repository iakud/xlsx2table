#ifndef _TABLE_STREAM_H_
#define _TABLE_STREAM_H_

#include <stdint.h>
#include <string>
#include <vector>
#include <map>
#include <iostream>

class TableStream {
public:
	TableStream(char* p);

	TableStream& operator>>(bool& b);
	TableStream& operator>>(int8_t& i);
	TableStream& operator>>(int16_t& i);
	TableStream& operator>>(int32_t & i);
	TableStream& operator>>(int64_t& i);
	TableStream& operator>>(uint8_t& n);
	TableStream& operator>>(uint16_t& n);
	TableStream& operator>>(uint32_t& n);
	TableStream& operator>>(uint64_t& n);
	TableStream& operator>>(float& f);
	TableStream& operator>>(double& f);
	TableStream& operator>>(std::string& s);
	template<class T>
	inline TableStream& operator>>(std::vector<T>& v);
	template<class K, class V>
	inline TableStream& operator>>(std::map<K, V>& m);

private:
	char* _p;
};

template<class T>
inline TableStream& TableStream::operator>>(std::vector<T>& v) {
	uint32_t size;
	v.clear();
	T t;
	*this >> size;
	for(uint32_t i = 0; i < size; ++i) {
		*this >> t;
		v.push_back(t);
	}
	return *this;
}

template<class K, class V>
inline TableStream& TableStream::operator>>(std::map<K, V>& m) {
	uint32_t size;
	m.clear();
	K k;
	V v;
	*this >> size;
	for(uint32_t i = 0; i < size; ++i) {
		*this >> k;
		*this >> v;
		m.insert(std::make_pair(k, v));
	}	
	return *this;
}

#endif // _TABLE_STREAM_H_
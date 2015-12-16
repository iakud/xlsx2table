#ifndef _TABLE_READER_H_
#define _TABLE_READER_H_

#include "TableStream.h"
#include <fstream> 
#include <stdint.h>
#include <set>

template<class Item>
class TableReader {
public:
	typedef std::set<Item*, typename Item::Compare> Set;

	TableReader() {}
	virtual ~TableReader() {}

	const Set& getItems() const;
	const Item* getItem(const typename Item::id_type& Id) const;

protected:
	Set _items;
	bool init(const char* filename);
};

template<class Item>
const typename TableReader<Item>::Set& TableReader<Item>::getItems() const {
	return _items;
}

template<class Item>
const Item* TableReader<Item>::getItem(const typename Item::id_type &id) const {
	Item item;
	item.setId(id);
	typename Set::const_iterator it = _items.find(&item);
	if(it != _items.end()) {
		return *it;
	} else {
		return NULL;
	}
}

template<class Item>
bool TableReader<Item>::init(const char* filename) {
	std::ifstream ifs(filename, std::ifstream::binary);
	if (!ifs.is_open())
		return false;

	ifs.seekg (0, std::ios::end);
	int length = static_cast<int>(ifs.tellg());
    ifs.seekg(0, std::ios::beg);
    char* buffer = new char[length];
    ifs.read(buffer, length);
    ifs.close();

    TableStream stream(buffer);
	uint32_t count;
	stream >> count;
	for (uint32_t i = 0; i < count; ++i) {
		Item* item = new Item();
		stream >> *item;
		_items.insert(item);
	}
	delete[] buffer;
	return true;
}

#endif // _TABLE_READER_H_
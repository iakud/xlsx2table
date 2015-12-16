#include "exampleTable.h"
#include <iostream>

int main() {
	exampleTable tb;
	const exampleTable::Set& items = tb.getItems();
	for (exampleTable::Set::const_iterator it = items.begin(); it != items.end(); ++it) {
		std::cout << "int16 Value:" << (*it)->i16 << std::endl;
		std::cout << "int32 Value:" << (*it)->i32 << std::endl;
		std::cout << "int64 Value:" << (*it)->i64 << std::endl;
		printf("float, double: %f, %f", (*it)->f, (*it)->df);
		std::cout << "string Value:" << (*it)->s << std::endl;
		int size = (*it)->vectorValue.size();
		for (int i = 0; i < size; ++i) {
			std::cout << "vector Value:" << (*it)->v[i] << std::endl;
		}
		std::map<std::string,std::string>::iterator itm = (*it)->m.begin();
		for (; itm!=(*it)->m.end(); ++itm) {
			std::cout << "map Value:" << itm->first << "," << itm->second << std::endl;
		}
	}
	return 0;
}
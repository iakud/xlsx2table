#include "exampleTable.h"
#include <iostream>

int main() {
	ExampleTable tb;
	const ExampleTable::Set& items = tb.getItems();
	for (ExampleTable::Set::const_iterator it = items.begin(); it != items.end(); ++it) {
		std::cout << "int16Value:" << (*it)->int16Value << std::endl;
		std::cout << "int32Value:" << (*it)->int32Value << std::endl;
		std::cout << "int64Value:" << (*it)->int64Value << std::endl;
		printf("%f,%f", (*it)->floatValue, (*it)->doubleValue);
		std::cout << "stringValue:" << (*it)->stringValue << std::endl;
		int size = (*it)->vectorValue.size();
		for (int i = 0; i < size; ++i) {
			std::cout << "vectorValue:" << (*it)->vectorValue[i] << std::endl;
		}
		std::map<std::string,std::string>::iterator itmap = (*it)->mapValue.begin();
		for (; itmap!=(*it)->mapValue.end(); ++itmap) {
			std::cout << "mapValue:" << itmap->first << "," << itmap->second << std::endl;
		}
	}
	return 0;
}
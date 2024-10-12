from typing import List, Optional
from pydantic import BaseModel, Field

class ArrayList(BaseModel):
    arr: List[int] = Field(default_factory=list)
    numOfData: int = 0
    curPosition: int = -1

    def LInsert(self, data: int) -> None:
        self.arr.append(data)
        self.numOfData += 1

    def LFirst(self) -> bool:
        if self.numOfData == 0:
            return False
        
        self.curPosition = 0
        return True
    
    def LNext(self) -> bool:
        if self.curPosition >= self.numOfData - 1:
            return False
        
        self.curPosition += 1
        return True
    
    def LRemove(self) -> Optional[int]:
        if self.curPosition == -1:
            return None
        
        rdata = self.arr[self.curPosition]
        del self.arr[self.curPosition]

        self.numOfData -= 1
        self.curPosition -= 1
        return rdata
    
    def LCount(self) -> int:
        return self.numOfData
        
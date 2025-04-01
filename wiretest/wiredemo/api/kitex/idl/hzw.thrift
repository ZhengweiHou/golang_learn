// Copyright 2021 CloudWeGo Authors 
// 
// Licensed under the Apache License, Version 2.0 (the "License"); 
// you may not use this file except in compliance with the License. 
// You may obtain a copy of the License at 
// 
//     http://www.apache.org/licenses/LICENSE-2.0 
// 
// Unless required by applicable law or agreed to in writing, software 
// distributed under the License is distributed on an "AS IS" BASIS, 
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. 
// See the License for the specific language governing permissions and 
// limitations under the License. 
// 

namespace go hzw

typedef i64 Timestamp

struct HzwDto {
    1: i64 Id
    2: string Name
    3: i32 Age
    4: i32 Version
    5: Timestamp CreatedAt
    6: Timestamp UpdatedAt
    7: Timestamp Time1
    8: Timestamp Time2
    9: Timestamp Time3
    10: double Decimal1
}

service HzwService {
    // 创建Hzw
    HzwDto CreateHzw(1: HzwDto hzwDto)
    
    // 根据ID查询Hzw
    HzwDto GetHzw(1: i64 id)
    
    // 创建Hzw事务测试
    HzwDto CreateHzwTxTest(1: HzwDto hzwDto)
}

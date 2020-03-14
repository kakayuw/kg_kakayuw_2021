# kg_kakayuw_2021
mapping question solution

|Author| Time | Language
|:--:|:--:|:--:|
|Hang Yu| 13/3/2020 | python / js / go
----
### Explanation
I use a hastable to solve this question. To check whether string `s` could be map to string `t`, it iterate through each position at `s` and `t` and try to create mapping from `char_s` to `chat_t`. If this mapping has been created before, it check whether the new mapping is same as the former one, otherwise the mapping mismatch and return False. While no mapping created before, it will create a mapping and continue parsing.

Time Complexity: O(N)

Space Complexity: O(N)
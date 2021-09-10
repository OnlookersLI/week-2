1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

   答：需要Wrap 这个 error给上层！当执行出现错误的时候要知道什么操作出现的sql.ErrNoRows,方便定位


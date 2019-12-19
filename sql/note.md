select * from tb1 where tb1_col1 in (select tb2_col1 from tb2 as s where tb2_col2 ='value');

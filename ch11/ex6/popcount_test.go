package popcount

import (
	"fmt"
	"testing"
)

var benchConf = struct {
	start, stop uint64
}{
	2e0,
	2e10,
}

func BenchmarkPopCountTable(b *testing.B) {
	for i := benchConf.start; i <= benchConf.stop; i *= 2 {
		i := i
		b.Run(fmt.Sprintf("x=%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				PopCountTable(i)
			}
		})
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := benchConf.start; i <= benchConf.stop; i *= 2 {
		i := i
		b.Run(fmt.Sprintf("x=%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				PopCountShift(i)
			}
		})
	}
}

func BenchmarkPopCountClearRightmostNonZeroBit(b *testing.B) {
	for i := benchConf.start; i <= benchConf.stop; i *= 2 {
		i := i
		b.Run(fmt.Sprintf("x=%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				PopCountClearRightmostNonZeroBit(i)
			}
		})
	}
}

// Benchmarks reveal that the table-based PopCount function is the winner. See results below:
//
// go test -bench=. -benchtime=1000x
// BenchmarkPopCountTable/x=2-8                  	    					1000	         	0.5840 ns/op
// BenchmarkPopCountTable/x=4-8                  	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=8-8                  	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=16-8                 	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=32-8                 	    					1000	         	0.5840 ns/op
// BenchmarkPopCountTable/x=64-8                 	    					1000	         	0.5410 ns/op
// BenchmarkPopCountTable/x=128-8                	    					1000	         	0.5410 ns/op
// BenchmarkPopCountTable/x=256-8                	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=512-8                	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=1024-8               	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=2048-8               	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=4096-8               	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=8192-8               	    					1000	         	0.5830 ns/op
// BenchmarkPopCountTable/x=16384-8              	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=32768-8              	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=65536-8              	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=131072-8             	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=262144-8             	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=524288-8             	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=1048576-8            	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=2097152-8            	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=4194304-8            	    					1000	         	0.5410 ns/op
// BenchmarkPopCountTable/x=8388608-8            	    					1000	         	0.6250 ns/op
// BenchmarkPopCountTable/x=16777216-8           	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=33554432-8           	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=67108864-8           	    					1000	         	0.5830 ns/op
// BenchmarkPopCountTable/x=134217728-8          	    					1000	         	0.5410 ns/op
// BenchmarkPopCountTable/x=268435456-8          	    					1000	         	0.4590 ns/op
// BenchmarkPopCountTable/x=536870912-8          	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=1073741824-8         	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=2147483648-8         	    					1000	         	0.4580 ns/op
// BenchmarkPopCountTable/x=4294967296-8         	    					1000	         	0.5420 ns/op
// BenchmarkPopCountTable/x=8589934592-8         	    					1000	         	0.5000 ns/op
// BenchmarkPopCountTable/x=17179869184-8        	    					1000	         	0.5830 ns/op
//
// BenchmarkPopCountShift/x=2-8                  	    					1000	        	29.00 ns/op
// BenchmarkPopCountShift/x=4-8                  	    					1000	        	32.75 ns/op
// BenchmarkPopCountShift/x=8-8                  	    					1000	        	32.67 ns/op
// BenchmarkPopCountShift/x=16-8                 	    					1000	        	32.83 ns/op
// BenchmarkPopCountShift/x=32-8                 	    					1000	        	29.29 ns/op
// BenchmarkPopCountShift/x=64-8                 	    					1000	        	32.46 ns/op
// BenchmarkPopCountShift/x=128-8                	    					1000	        	28.46 ns/op
// BenchmarkPopCountShift/x=256-8                	    					1000	        	28.29 ns/op
// BenchmarkPopCountShift/x=512-8                	    					1000	        	28.38 ns/op
// BenchmarkPopCountShift/x=1024-8               	    					1000	        	29.00 ns/op
// BenchmarkPopCountShift/x=2048-8               	    					1000	        	32.33 ns/op
// BenchmarkPopCountShift/x=4096-8               	    					1000	        	37.96 ns/op
// BenchmarkPopCountShift/x=8192-8               	    					1000	        	28.96 ns/op
// BenchmarkPopCountShift/x=16384-8              	    					1000	        	28.46 ns/op
// BenchmarkPopCountShift/x=32768-8              	    					1000	        	27.67 ns/op
// BenchmarkPopCountShift/x=65536-8              	    					1000	        	28.54 ns/op
// BenchmarkPopCountShift/x=131072-8             	    					1000	        	28.33 ns/op
// BenchmarkPopCountShift/x=262144-8             	    					1000	        	27.75 ns/op
// BenchmarkPopCountShift/x=524288-8             	    					1000	        	32.75 ns/op
// BenchmarkPopCountShift/x=1048576-8            	    					1000	        	28.00 ns/op
// BenchmarkPopCountShift/x=2097152-8            	    					1000	        	28.62 ns/op
// BenchmarkPopCountShift/x=4194304-8            	    					1000	        	31.83 ns/op
// BenchmarkPopCountShift/x=8388608-8            	    					1000	        	28.29 ns/op
// BenchmarkPopCountShift/x=16777216-8           	    					1000	        	32.33 ns/op
// BenchmarkPopCountShift/x=33554432-8           	    					1000	        	28.38 ns/op
// BenchmarkPopCountShift/x=67108864-8           	    					1000	        	28.71 ns/op
// BenchmarkPopCountShift/x=134217728-8          	    					1000	        	32.00 ns/op
// BenchmarkPopCountShift/x=268435456-8          	    					1000	        	32.92 ns/op
// BenchmarkPopCountShift/x=536870912-8          	    					1000	        	32.92 ns/op
// BenchmarkPopCountShift/x=1073741824-8         	    					1000	        	28.50 ns/op
// BenchmarkPopCountShift/x=2147483648-8         	    					1000	        	29.00 ns/op
// BenchmarkPopCountShift/x=4294967296-8         	    					1000	        	27.67 ns/op
// BenchmarkPopCountShift/x=8589934592-8         	    					1000	        	32.50 ns/op
// BenchmarkPopCountShift/x=17179869184-8        	    					1000	        	29.08 ns/op
//
// BenchmarkPopCountClearRightmostNonZeroBit/x=2-8         	    			1000	        	23.67 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=4-8         	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=8-8         	    			1000	        	23.25 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=16-8        	    			1000	        	22.92 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=32-8        	    			1000	        	23.17 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=64-8        	    			1000	        	23.38 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=128-8       	    			1000	        	23.12 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=256-8       	    			1000	        	23.46 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=512-8       	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=1024-8      	    			1000	        	23.92 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=2048-8      	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=4096-8      	    			1000	        	22.92 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=8192-8      	    			1000	        	22.92 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=16384-8     	    			1000	        	22.88 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=32768-8     	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=65536-8     	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=131072-8    	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=262144-8    	    			1000	        	23.17 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=524288-8    	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=1048576-8   	    			1000	        	23.12 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=2097152-8   	    			1000	        	23.12 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=4194304-8   	    			1000	        	23.00 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=8388608-8   	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=16777216-8  	    			1000	        	23.08 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=33554432-8  	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=67108864-8  	    			1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=134217728-8 	    			1000	        	23.29 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=268435456-8 	    			1000	        	23.62 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=536870912-8 	    			1000	        	23.33 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=1073741824-8         	    1000	        	22.96 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=2147483648-8         	    1000	        	23.12 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=4294967296-8         	    1000	        	23.71 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=8589934592-8         	    1000	        	23.12 ns/op
// BenchmarkPopCountClearRightmostNonZeroBit/x=17179869184-8        	    1000	        	23.08 ns/op
//
// These benchmarks, though, assume that the pre-computed table has been initialized beforehand,
// but it's fine to assume it as the table would be initialized just once in real-world code.
//
// Also, PopCountTable breaks even when b.N >= 100,000,000. See results below:
//
// go test -bench=PopCountTable -benchtime=1000000x
// BenchmarkPopCountTable/x=2-8         	         1000000	         0.3349 ns/op
// BenchmarkPopCountTable/x=4-8         	         1000000	         0.3349 ns/op
// BenchmarkPopCountTable/x=8-8         	         1000000	         0.3350 ns/op
// BenchmarkPopCountTable/x=16-8        	         1000000	         0.3350 ns/op
// BenchmarkPopCountTable/x=32-8        	         1000000	         0.3259 ns/op
// BenchmarkPopCountTable/x=64-8        	         1000000	         0.3237 ns/op
// BenchmarkPopCountTable/x=128-8       	         1000000	         0.3232 ns/op
// BenchmarkPopCountTable/x=256-8       	         1000000	         0.3247 ns/op
// BenchmarkPopCountTable/x=512-8       	         1000000	         0.3248 ns/op
// BenchmarkPopCountTable/x=1024-8      	         1000000	         0.3251 ns/op
// BenchmarkPopCountTable/x=2048-8      	         1000000	         0.3252 ns/op
// BenchmarkPopCountTable/x=4096-8      	         1000000	         0.3239 ns/op
// BenchmarkPopCountTable/x=8192-8      	         1000000	         0.3242 ns/op
// BenchmarkPopCountTable/x=16384-8     	         1000000	         0.3224 ns/op
// BenchmarkPopCountTable/x=32768-8     	         1000000	         0.3200 ns/op
// BenchmarkPopCountTable/x=65536-8     	         1000000	         0.3196 ns/op
// BenchmarkPopCountTable/x=131072-8    	         1000000	         0.3208 ns/op
// BenchmarkPopCountTable/x=262144-8    	         1000000	         0.3227 ns/op
// BenchmarkPopCountTable/x=524288-8    	         1000000	         0.3211 ns/op
// BenchmarkPopCountTable/x=1048576-8   	         1000000	         0.3203 ns/op
// BenchmarkPopCountTable/x=2097152-8   	         1000000	         0.3230 ns/op
// BenchmarkPopCountTable/x=4194304-8   	         1000000	         0.3198 ns/op
// BenchmarkPopCountTable/x=8388608-8   	         1000000	         0.3220 ns/op
// BenchmarkPopCountTable/x=16777216-8  	         1000000	         0.3188 ns/op
// BenchmarkPopCountTable/x=33554432-8  	         1000000	         0.3261 ns/op
// BenchmarkPopCountTable/x=67108864-8  	         1000000	         0.3476 ns/op
// BenchmarkPopCountTable/x=134217728-8 	         1000000	         0.3385 ns/op
// BenchmarkPopCountTable/x=268435456-8 	         1000000	         0.3381 ns/op
// BenchmarkPopCountTable/x=536870912-8 	         1000000	         0.3255 ns/op
// BenchmarkPopCountTable/x=1073741824-8         	 1000000	         0.3374 ns/op
// BenchmarkPopCountTable/x=2147483648-8         	 1000000	         0.3245 ns/op
// BenchmarkPopCountTable/x=4294967296-8         	 1000000	         0.3276 ns/op
// BenchmarkPopCountTable/x=8589934592-8         	 1000000	         0.3255 ns/op
// BenchmarkPopCountTable/x=17179869184-8        	 1000000	         0.3255 ns/op
//
// go test -bench=PopCountTable -benchtime=10000000x
// BenchmarkPopCountTable/x=2-8         	        10000000	         0.3347 ns/op
// BenchmarkPopCountTable/x=4-8         	        10000000	         0.3301 ns/op
// BenchmarkPopCountTable/x=8-8         	        10000000	         0.3232 ns/op
// BenchmarkPopCountTable/x=16-8        	        10000000	         0.3192 ns/op
// BenchmarkPopCountTable/x=32-8        	        10000000	         0.3182 ns/op
// BenchmarkPopCountTable/x=64-8        	        10000000	         0.3193 ns/op
// BenchmarkPopCountTable/x=128-8       	        10000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=256-8       	        10000000	         0.3126 ns/op
// BenchmarkPopCountTable/x=512-8       	        10000000	         0.3127 ns/op
// BenchmarkPopCountTable/x=1024-8      	        10000000	         0.3376 ns/op
// BenchmarkPopCountTable/x=2048-8      	        10000000	         0.3260 ns/op
// BenchmarkPopCountTable/x=4096-8      	        10000000	         0.3166 ns/op
// BenchmarkPopCountTable/x=8192-8      	        10000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=16384-8     	        10000000	         0.3127 ns/op
// BenchmarkPopCountTable/x=32768-8     	        10000000	         0.3200 ns/op
// BenchmarkPopCountTable/x=65536-8     	        10000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=131072-8    	        10000000	         0.3125 ns/op
// BenchmarkPopCountTable/x=262144-8    	        10000000	         0.3133 ns/op
// BenchmarkPopCountTable/x=524288-8    	        10000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=1048576-8   	        10000000	         0.3161 ns/op
// BenchmarkPopCountTable/x=2097152-8   	        10000000	         0.3347 ns/op
// BenchmarkPopCountTable/x=4194304-8   	        10000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=8388608-8   	        10000000	         0.3137 ns/op
// BenchmarkPopCountTable/x=16777216-8  	        10000000	         0.3141 ns/op
// BenchmarkPopCountTable/x=33554432-8  	        10000000	         0.3146 ns/op
// BenchmarkPopCountTable/x=67108864-8  	        10000000	         0.3148 ns/op
// BenchmarkPopCountTable/x=134217728-8 	        10000000	         0.3125 ns/op
// BenchmarkPopCountTable/x=268435456-8 	        10000000	         0.3126 ns/op
// BenchmarkPopCountTable/x=536870912-8 	        10000000	         0.3138 ns/op
// BenchmarkPopCountTable/x=1073741824-8         	10000000	         0.3152 ns/op
// BenchmarkPopCountTable/x=2147483648-8         	10000000	         0.3178 ns/op
// BenchmarkPopCountTable/x=4294967296-8         	10000000	         0.3311 ns/op
// BenchmarkPopCountTable/x=8589934592-8         	10000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=17179869184-8        	10000000	         0.3134 ns/op
//
// go test -bench=PopCountTable -benchtime=100000000x
// BenchmarkPopCountTable/x=2-8         			100000000	         0.3178 ns/op
// BenchmarkPopCountTable/x=4-8         			100000000	         0.3154 ns/op
// BenchmarkPopCountTable/x=8-8         			100000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=16-8        			100000000	         0.3145 ns/op
// BenchmarkPopCountTable/x=32-8        			100000000	         0.3126 ns/op
// BenchmarkPopCountTable/x=64-8        			100000000	         0.3150 ns/op
// BenchmarkPopCountTable/x=128-8       			100000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=256-8       			100000000	         0.3167 ns/op
// BenchmarkPopCountTable/x=512-8       			100000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=1024-8      			100000000	         0.3147 ns/op
// BenchmarkPopCountTable/x=2048-8      			100000000	         0.3133 ns/op
// BenchmarkPopCountTable/x=4096-8      			100000000	         0.3145 ns/op
// BenchmarkPopCountTable/x=8192-8      			100000000	         0.3129 ns/op
// BenchmarkPopCountTable/x=16384-8     			100000000	         0.3173 ns/op
// BenchmarkPopCountTable/x=32768-8     			100000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=65536-8     			100000000	         0.3150 ns/op
// BenchmarkPopCountTable/x=131072-8    			100000000	         0.3129 ns/op
// BenchmarkPopCountTable/x=262144-8    			100000000	         0.3147 ns/op
// BenchmarkPopCountTable/x=524288-8    			100000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=1048576-8   			100000000	         0.3146 ns/op
// BenchmarkPopCountTable/x=2097152-8   			100000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=4194304-8   			100000000	         0.3146 ns/op
// BenchmarkPopCountTable/x=8388608-8   			100000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=16777216-8  			100000000	         0.3153 ns/op
// BenchmarkPopCountTable/x=33554432-8  			100000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=67108864-8  			100000000	         0.3146 ns/op
// BenchmarkPopCountTable/x=134217728-8 			100000000	         0.3128 ns/op
// BenchmarkPopCountTable/x=268435456-8 			100000000	         0.3149 ns/op
// BenchmarkPopCountTable/x=536870912-8 			100000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=1073741824-8         	100000000	         0.3147 ns/op
// BenchmarkPopCountTable/x=2147483648-8         	100000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=4294967296-8         	100000000	         0.3149 ns/op
// BenchmarkPopCountTable/x=8589934592-8         	100000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=17179869184-8        	100000000	         0.3148 ns/op
//
// go test -bench=PopCountTable -benchtime=1000000000x
// BenchmarkPopCountTable/x=2-8         			1000000000	         0.3138 ns/op
// BenchmarkPopCountTable/x=4-8         			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=8-8         			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=16-8        			1000000000	         0.3156 ns/op
// BenchmarkPopCountTable/x=32-8        			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=64-8        			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=128-8       			1000000000	         0.3138 ns/op
// BenchmarkPopCountTable/x=256-8       			1000000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=512-8       			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=1024-8      			1000000000	         0.3133 ns/op
// BenchmarkPopCountTable/x=2048-8      			1000000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=4096-8      			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=8192-8      			1000000000	         0.3133 ns/op
// BenchmarkPopCountTable/x=16384-8     			1000000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=32768-8     			1000000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=65536-8     			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=131072-8    			1000000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=262144-8    			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=524288-8    			1000000000	         0.3129 ns/op
// BenchmarkPopCountTable/x=1048576-8   			1000000000	         0.3135 ns/op
// BenchmarkPopCountTable/x=2097152-8   			1000000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=4194304-8   			1000000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=8388608-8   			1000000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=16777216-8  			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=33554432-8  			1000000000	         0.3130 ns/op
// BenchmarkPopCountTable/x=67108864-8  			1000000000	         0.3136 ns/op
// BenchmarkPopCountTable/x=134217728-8 			1000000000	         0.3142 ns/op
// BenchmarkPopCountTable/x=268435456-8 			1000000000	         0.3131 ns/op
// BenchmarkPopCountTable/x=536870912-8 			1000000000	         0.3136 ns/op
// BenchmarkPopCountTable/x=1073741824-8         	1000000000	         0.3136 ns/op
// BenchmarkPopCountTable/x=2147483648-8         	1000000000	         0.3129 ns/op
// BenchmarkPopCountTable/x=4294967296-8         	1000000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=8589934592-8         	1000000000	         0.3132 ns/op
// BenchmarkPopCountTable/x=17179869184-8        	1000000000	         0.3130 ns/op

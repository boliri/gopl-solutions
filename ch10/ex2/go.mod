module archiveread

go 1.19

replace (
	arch => ./arch
	arch/tar => ./arch/tar
	arch/zip => ./arch/zip
)

require (
	arch v0.0.0
	arch/tar v0.0.0
	arch/zip v0.0.0
)

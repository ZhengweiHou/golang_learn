#!/bin/bash
targetdir=hzwent

rm -rf  schema
mkdir schema
#cp templ/* schema/

#cat templ/student.temp > schema/student.go
ls templ | while read line; do echo "${line}"; cp templ/${line} ./schema/${line%.*}.go  ; done

rm -rf ${targetdir}

mkdir ${targetdir}
echo "package ${targetdir}" > ${targetdir}/dir.go



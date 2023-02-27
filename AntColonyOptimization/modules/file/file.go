package file

import (
  "os"
  "strings"
  "bufio"
  "io"
  "strconv"
)


func LoadCityFile(fileName string) [][]int{
  fp, err := os.Open(fileName)
  
  if err != nil{
    panic(err)
  }
  defer fp.Close()

  reader := bufio.NewReaderSize(fp, 64)
  var tmp []byte
  var dim int
  var cityCoordinate [][]int
  startReadingCoordinate := false
  currentCoordinateIndex := -1

  for{
    line, isPrefix, err := reader.ReadLine()
    if err == io.EOF{
      break
    }
    if err != nil{
      panic(err)
    }
    tmp = append(tmp, line...)

    if !isPrefix{
      l := string(tmp)
      if l == "EOF"{
        break
      }
      tmp = nil

      if strings.Index(l, "DIMENSION") != -1{
        splittedDimensionLine := strings.Split(l, ": ")
        dim, _ = strconv.Atoi(splittedDimensionLine[1])

        for i := 0; i<dim; i++{
          cityCoordinate = append(cityCoordinate, make([]int, 2)) 
        }
        continue
      }

      if strings.Index(l, "NODE_COORD_SECTION") != -1{
        startReadingCoordinate = true
        continue
      }

      if startReadingCoordinate{
        splittedCoordinateLine := strings.Split(l, " ")
        xCoordinate, _ := strconv.Atoi(splittedCoordinateLine[1])
        yCoordinate, _ := strconv.Atoi(splittedCoordinateLine[2])
        currentCoordinateIndex++
        cityCoordinate[currentCoordinateIndex][0] = xCoordinate
        cityCoordinate[currentCoordinateIndex][1] = yCoordinate

      }
    }
  }
  return cityCoordinate
}

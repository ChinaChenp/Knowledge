/* fread example: read an entire file */
#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <vector>

using namespace std;

int readAllFile (const std::string &file_name, std::vector<char> & out) {
  FILE *p = fopen (file_name.c_str() , "rb" );
  if (p == NULL) {
      return -1;
  }

  // 文件大小
  fseek (p , 0 , SEEK_END);
  int file_size = ftell (p);
  rewind (p);

  out.resize(file_size);
  int result = fread (&out[0], 1, file_size, p);
  if (result != file_size) {
      return -1;
  }

  fclose (p);
  return 0;
}

int readAllFileLineByLine(const std::string &file_name, std::vector<std::string> & out) {
  FILE *fp = fopen(file_name.c_str(), "rb");
  if (fp == NULL) {
      return -1;
  }

  size_t total = 0;
  int len = 0;
  char *line_buf = NULL;
  while((len = ::getline(&line_buf, &total, fp)) > 0) {
      line_buf[len-1] = '\0';
      out.emplace_back(line_buf);
  }

  fclose(fp);
  free(line_buf);
  return 0;
}

int main(int argc, char *argv[]) {
    std::string file_name(argv[1]);

    std::vector<std::string> out;
    clock_t start = clock();
    int err = readAllFileLineByLine(file_name, out);
    //int err = readAllFile(file_name, out);
    clock_t end = clock();
    double pro = (double)(end-start)/CLOCKS_PER_SEC * 1000;
    if (err != 0) {
        printf("error");
    } else {
        for (auto & s : out) {
            printf("out[%lu]: %s--\n", s.size(), s.c_str());
        }
        //printf("out : %s\n", string(out.begin(), out.end()).c_str());
    }
    printf("total: %f\n", pro);
    return 0;
}
